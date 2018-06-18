package mexgen

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/mobiledgex/edge-cloud/gensupport"
	"github.com/mobiledgex/edge-cloud/protogen"
)

func RegisterMex() {
	generator.RegisterPlugin(new(mex))
}

func init() {
	generator.RegisterPlugin(new(mex))
}

type mex struct {
	gen           *generator.Generator
	msgs          map[string]*descriptor.DescriptorProto
	cudTemplate   *template.Template
	importUtil    bool
	importStrings bool
	support       gensupport.PluginSupport
}

func (m *mex) Name() string {
	return "mex"
}

func (m *mex) Init(gen *generator.Generator) {
	m.gen = gen
	m.msgs = make(map[string]*descriptor.DescriptorProto)
	m.cudTemplate = template.Must(template.New("cud").Parse(cudTemplateIn))
	m.support.Init(nil)
}

// P forwards to g.gen.P
func (m *mex) P(args ...interface{}) {
	m.gen.P(args...)
}

func (m *mex) Generate(file *generator.FileDescriptor) {
	m.support.InitFile()
	m.support.SetPbGoPackage(file.GetPackage())
	m.importUtil = false
	m.importStrings = false
	for _, desc := range file.Messages() {
		m.generateMessage(file, desc)
	}
	if len(file.FileDescriptorProto.Service) != 0 {
		for _, service := range file.FileDescriptorProto.Service {
			m.generateService(file, service)
		}
	}
}

func (m *mex) GenerateImports(file *generator.FileDescriptor) {
	hasGenerateCud := false
	hasGenerateCache := false
	for _, desc := range file.Messages() {
		msg := desc.DescriptorProto
		if GetGenerateCud(msg) {
			hasGenerateCud = true
			if GetGenerateCache(msg) {
				hasGenerateCache = true
			}
		}
		m.msgs[*msg.Name] = msg
	}
	if hasGenerateCud {
		m.gen.PrintImport("", "encoding/json")
		m.gen.PrintImport("", "github.com/mobiledgex/edge-cloud/objstore")
	}
	if hasGenerateCache {
		m.gen.PrintImport("", "sync")
	}
	if m.importUtil {
		m.gen.PrintImport("", "github.com/mobiledgex/edge-cloud/util")
	}
	if m.importStrings {
		m.gen.PrintImport("strings", "strings")
	}
	m.support.PrintUsedImports(m.gen)
}

func (m *mex) generateFieldMatches(message *descriptor.DescriptorProto, field *descriptor.FieldDescriptorProto) {
	if field.Type == nil {
		return
	}
	if *field.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
		// TODO: matches support for repeated fields
		return
	}
	name := generator.CamelCase(*field.Name)
	switch *field.Type {
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		nullcheck := ""
		ref := "&"
		if gogoproto.IsNullable(field) {
			nullcheck = fmt.Sprintf("filter.%s != nil && m.%s != nil && ", name, name)
			ref = ""
		}
		subDesc := gensupport.GetDesc(m.gen, field.GetTypeName())
		printedCheck := true
		if *field.TypeName == ".google.protobuf.Timestamp" {
			m.P("if ", nullcheck, "(m.", name, ".Seconds != filter.", name, ".Seconds || m.", name, ".Nanos != filter.", name, ".Nanos) {")
		} else if GetGenerateMatches(subDesc.DescriptorProto) {
			m.P("if ", nullcheck, "!m.", name, ".Matches(", ref, "filter.", name, ") {")
		} else {
			printedCheck = false
		}
		if printedCheck {
			m.P("return false")
			m.P("}")
		}
	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		// deprecated in proto3
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		// TODO
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		m.P("if filter.", name, " != \"\" && filter.", name, " != m.", name, "{")
		m.P("return false")
		m.P("}")
	default:
		m.P("if filter.", name, " != 0 && filter.", name, " != m.", name, "{")
		m.P("return false")
		m.P("}")
	}
}

func (m *mex) printCopyInMakeArray(name string, desc *generator.Descriptor, field *descriptor.FieldDescriptorProto) {
	typ, _ := m.gen.GoType(desc, field)
	m.P("if m.", name, " == nil || len(m.", name, ") < len(src.", name, ") {")
	m.P("m.", name, " = make(", typ, ", len(src.", name, "))")
	m.P("}")
}

func (m *mex) getFieldDesc(field *descriptor.FieldDescriptorProto) *generator.Descriptor {
	obj := m.gen.ObjectNamed(field.GetTypeName())
	if obj == nil {
		return nil
	}
	desc, ok := obj.(*generator.Descriptor)
	if ok {
		return desc
	}
	return nil
}

func (m *mex) generateFields(names, nums []string, desc *generator.Descriptor) {
	message := desc.DescriptorProto
	for ii, field := range message.Field {
		if ii == 0 && *field.Name == "fields" {
			continue
		}
		name := generator.CamelCase(*field.Name)
		num := fmt.Sprintf("%d", *field.Number)
		switch *field.Type {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			subDesc := gensupport.GetDesc(m.gen, field.GetTypeName())
			m.generateFields(append(names, name), append(nums, num), subDesc)
		default:
			m.P("const ", strings.Join(append(names, name), ""), " = \"", strings.Join(append(nums, num), "."), "\"")
		}
	}
}

func (m *mex) generateAllFields(names, nums []string, desc *generator.Descriptor) {
	message := desc.DescriptorProto
	for ii, field := range message.Field {
		if ii == 0 && *field.Name == "fields" {
			continue
		}
		name := generator.CamelCase(*field.Name)
		num := fmt.Sprintf("%d", *field.Number)
		switch *field.Type {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			subDesc := gensupport.GetDesc(m.gen, field.GetTypeName())
			m.generateAllFields(append(names, name), append(nums, num), subDesc)
		default:
			m.P(strings.Join(append(names, name), ""), ",")
		}
	}
}

func (m *mex) generateCopyIn(parents, nums []string, desc *generator.Descriptor, visited []*generator.Descriptor, hasGrpcFields bool) {
	if gensupport.WasVisited(desc, visited) {
		return
	}
	for ii, field := range desc.DescriptorProto.Field {
		if ii == 0 && *field.Name == "fields" {
			continue
		}
		if field.OneofIndex != nil {
			// no support for copy OneOf fields
			continue
		}

		name := generator.CamelCase(*field.Name)
		hierName := strings.Join(append(parents, name), ".")
		num := fmt.Sprintf("%d", *field.Number)
		idx := ""
		nullableMessage := false
		if *field.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE && gogoproto.IsNullable(field) {
			nullableMessage = true
		}

		if hasGrpcFields {
			numStr := strings.Join(append(nums, num), ".")
			nilCheck := ""
			if nullableMessage {
				nilCheck = " && src." + hierName + " != nil"
			}
			m.P("if _, set := fmap[\"", numStr, "\"]; set", nilCheck, " {")
		} else if nullableMessage {
			m.P("if src.", hierName, " != nil {")
		}
		if *field.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			m.printCopyInMakeArray(hierName, desc, field)
			if *field.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
				incr := fmt.Sprintf("i%d", len(parents))
				m.P("for ", incr, " := 0; ", incr, " < len(src.", hierName, "); ", incr, "++ {")
				idx = "[" + incr + "]"
			}
		}
		switch *field.Type {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			subDesc := gensupport.GetDesc(m.gen, field.GetTypeName())
			if gogoproto.IsNullable(field) {
				typ := m.support.FQTypeName(m.gen, subDesc)
				m.P("m.", hierName, idx, " = &", typ, "{}")
			}
			m.generateCopyIn(append(parents, name+idx), append(nums, num), subDesc, append(visited, desc), hasGrpcFields)
		case descriptor.FieldDescriptorProto_TYPE_GROUP:
			// deprecated in proto3
		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			m.printCopyInMakeArray(hierName, desc, field)
			m.P("copy(m.", hierName, ", src.", hierName, ")")
		default:
			if *field.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED {
				m.P("copy(m.", hierName, ", src.", hierName, ")")
			} else {
				m.P("m.", hierName, " = src.", hierName)
			}
		}
		if *field.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED && *field.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			m.P("}")
		}
		if hasGrpcFields || nullableMessage {
			m.P("}")
		}
	}
}

type cudTemplateArgs struct {
	Name      string
	KeyName   string
	CudName   string
	HasFields bool
	GenCache  bool
}

var cudTemplateIn = `
func (s *{{.Name}}) HasFields() bool {
{{- if (.HasFields)}}
	return true
{{- else}}
	return false
{{- end}}
}

type {{.Name}}Store struct {
	objstore objstore.ObjStore
	list{{.Name}} map[{{.Name}}Key]struct{}
}

func New{{.Name}}Store(objstore objstore.ObjStore) {{.Name}}Store {
	return {{.Name}}Store{objstore: objstore}
}

type {{.Name}}Cacher interface {
	Sync{{.Name}}Update(m *{{.Name}}, rev int64)
	Sync{{.Name}}Delete(m *{{.Name}}, rev int64)
	Sync{{.Name}}Prune(current map[{{.Name}}Key]struct{})
	Sync{{.Name}}RevOnly(rev int64)
}

func (s *{{.Name}}Store) Create(m *{{.Name}}, wait func(int64)) (*Result, error) {
	err := m.Validate()
	if err != nil { return nil, err }
	key := objstore.DbKeyString(m.GetKey())
	val, err := json.Marshal(m)
	if err != nil { return nil, err }
	rev, err := s.objstore.Create(key, string(val))
	if err != nil { return nil, err }
	if wait != nil {
		wait(rev)
	}
	return &Result{}, err
}

func (s *{{.Name}}Store) Update(m *{{.Name}}, wait func(int64)) (*Result, error) {
	err := m.Validate()
	if err != nil { return nil, err }
	key := objstore.DbKeyString(m.GetKey())
	var vers int64 = 0
{{- if (.HasFields)}}
	curBytes, vers, err := s.objstore.Get(key)
	if err != nil { return nil, err }
	var cur {{.Name}}
	err = json.Unmarshal(curBytes, &cur)
	if err != nil { return nil, err }
	cur.CopyInFields(m)
	// never save fields
	cur.Fields = nil
	val, err := json.Marshal(cur)
{{- else}}
	val, err := json.Marshal(m)
{{- end}}
	if err != nil { return nil, err }
	rev, err := s.objstore.Update(key, string(val), vers)
	if err != nil { return nil, err }
	if wait != nil {
		wait(rev)
	}
	return &Result{}, err
}

func (s *{{.Name}}Store) Delete(m *{{.Name}}, wait func(int64)) (*Result, error) {
	err := m.GetKey().Validate()
	if err != nil { return nil, err }
	key := objstore.DbKeyString(m.GetKey())
	rev, err := s.objstore.Delete(key)
	if err != nil { return nil, err }
	if wait != nil {
		wait(rev)
	}
	return &Result{}, err
}

type {{.Name}}Cb func(m *{{.Name}}) error

func (s *{{.Name}}Store) LoadAll(cb {{.Name}}Cb) error {
	loadkey := objstore.DbKeyPrefixString(&{{.Name}}Key{})
	err := s.objstore.List(loadkey, func(key, val []byte, rev int64) error {
		var obj {{.Name}}
		err := json.Unmarshal(val, &obj)
		if err != nil {
			util.WarnLog("Failed to parse {{.Name}} data", "val", string(val))
			return nil
		}
		err = cb(&obj)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *{{.Name}}Store) LoadOne(key string) (*{{.Name}}, int64, error) {
	val, rev, err := s.objstore.Get(key)
	if err != nil {
		return nil, 0, err
	}
	var obj {{.Name}}
	err = json.Unmarshal(val, &obj)
	if err != nil {
		util.DebugLog(util.DebugLevelApi, "Failed to parse {{.Name}} data", "val", string(val))
		return nil, 0, err
	}
	return &obj, rev, nil
}

// Sync will sync changes for any {{.Name}} objects.
func (s *{{.Name}}Store) Sync(ctx context.Context, cacher {{.Name}}Cacher) error {
	str := objstore.DbKeyPrefixString(&{{.Name}}Key{}) + "/"
	return s.objstore.Sync(ctx, str, func(in *objstore.SyncCbData) {
		obj := {{.Name}}{}
		// Even on parse error, we should still call back to keep
		// the revision numbers in sync so no caller hangs on wait.
		action := in.Action
		if action == objstore.SyncUpdate || action == objstore.SyncList {
			err := json.Unmarshal(in.Value, &obj)
			if err != nil {
				util.WarnLog("Failed to parse {{.Name}} data", "val", string(in.Value))
				action = objstore.SyncRevOnly
			}
		} else if action == objstore.SyncDelete {
			keystr := objstore.DbKeyPrefixRemove(string(in.Key))
			{{.Name}}KeyStringParse(keystr, obj.GetKey())
		}
		util.DebugLog(util.DebugLevelApi, "Sync cb", "action", objstore.SyncActionStrs[in.Action], "key", string(in.Key), "value", string(in.Value), "rev", in.Rev)
		switch action {
		case objstore.SyncUpdate:
			cacher.Sync{{.Name}}Update(&obj, in.Rev)
		case objstore.SyncDelete:
			cacher.Sync{{.Name}}Delete(&obj, in.Rev)
		case objstore.SyncListStart:
			s.list{{.Name}} = make(map[{{.Name}}Key]struct{})
		case objstore.SyncList:
			s.list{{.Name}}[obj.Key] = struct{}{}
			cacher.Sync{{.Name}}Update(&obj, in.Rev)
		case objstore.SyncListEnd:
			cacher.Sync{{.Name}}Prune(s.list{{.Name}})
			s.list{{.Name}} = nil
		case objstore.SyncRevOnly:
			cacher.Sync{{.Name}}RevOnly(in.Rev)
		}
	})
}

{{if (.GenCache)}}
// {{.Name}}Cache caches {{.Name}} objects in memory in a hash table
// and keeps them in sync with the database.
type {{.Name}}Cache struct {
	Store *{{.Name}}Store
	Objs map[{{.Name}}Key]*{{.Name}}
	Rev int64
	Mux util.Mutex
	Cond sync.Cond
	initWait bool
	syncDone bool
	syncCancel context.CancelFunc
	notifyCb func(obj *{{.Name}}Key)
}

func New{{.Name}}Cache(store *{{.Name}}Store) *{{.Name}}Cache {
	cache := {{.Name}}Cache{
		Store: store,
		Objs: make(map[{{.Name}}Key]*{{.Name}}),
		initWait: true,
	}
	cache.Mux.InitCond(&cache.Cond)

	ctx, cancel := context.WithCancel(context.Background())
	cache.syncCancel = cancel
	go func() {
		err := cache.Store.Sync(ctx, &cache)
		if err != nil {
			util.WarnLog("{{.Name}} Sync failed", "err", err)
		}
		cache.syncDone = true
		cache.Cond.Broadcast()
	}()
	return &cache
}

func (c *{{.Name}}Cache) WaitInitSyncDone() {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	for c.initWait {
		c.Cond.Wait()
	}
}

func (c *{{.Name}}Cache) Done() {
	c.syncCancel()
	c.Mux.Lock()
	defer c.Mux.Unlock()
	for !c.syncDone {
		c.Cond.Wait()
	}
}

func (c *{{.Name}}Cache) Get(key *{{.Name}}Key, valbuf *{{.Name}}) bool {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	inst, found := c.Objs[*key]
	if found {
		*valbuf = *inst
	}
	return found
}

func (c *{{.Name}}Cache) HasKey(key *{{.Name}}Key) bool {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	_, found := c.Objs[*key]
	return found
}

func (c *{{.Name}}Cache) GetAllKeys(keys map[{{.Name}}Key]struct{}) {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	for key, _ := range c.Objs {
		keys[key] = struct{}{}
	}
}

func (c *{{.Name}}Cache) Sync{{.Name}}Update(in *{{.Name}}, rev int64) {
	c.Mux.Lock()
	c.Objs[*in.GetKey()] = in
	c.Rev = rev
	util.DebugLog(util.DebugLevelApi, "SyncUpdate", "obj", in, "rev", rev)
	c.Cond.Broadcast()
	c.Mux.Unlock()
	if c.notifyCb != nil {
		c.notifyCb(in.GetKey())
	}
}

func (c *{{.Name}}Cache) Sync{{.Name}}Delete(in *{{.Name}}, rev int64) {
	c.Mux.Lock()
	delete(c.Objs, *in.GetKey())
	c.Rev = rev
	util.DebugLog(util.DebugLevelApi, "SyncUpdate", "key", in.GetKey(), "rev", rev)
	c.Cond.Broadcast()
	c.Mux.Unlock()
	if c.notifyCb != nil {
		c.notifyCb(in.GetKey())
	}
}

func (c *{{.Name}}Cache) Sync{{.Name}}Prune(current map[{{.Name}}Key]struct{}) {
	deleted := make(map[{{.Name}}Key]struct{})
	c.Mux.Lock()
	for key, _ := range c.Objs {
		if _, found := current[key]; !found {
			delete(c.Objs, key)
			deleted[key] = struct{}{}
		}
	}
	if c.initWait {
		c.initWait = false
		c.Cond.Broadcast()
	}
	c.Mux.Unlock()
	if c.notifyCb != nil {
		for key, _ := range deleted {
			c.notifyCb(&key)
		}
	}
}

func (c *{{.Name}}Cache) Sync{{.Name}}RevOnly(rev int64) {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	c.Rev = rev
	util.DebugLog(util.DebugLevelApi, "SyncRevOnly", "rev", rev)
	c.Cond.Broadcast()
}

func (c *{{.Name}}Cache) SyncWait(rev int64) {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	util.DebugLog(util.DebugLevelApi, "SyncWait", "cache-rev", c.Rev, "wait-rev", rev)
	for c.Rev < rev {
		c.Cond.Wait()
	}
}

func (c *{{.Name}}Cache) Show(filter *{{.Name}}, cb func(ret *{{.Name}}) error) error {
	util.DebugLog(util.DebugLevelApi, "Show {{.Name}}", "count", len(c.Objs))
	c.Mux.Lock()
	defer c.Mux.Unlock()
	for _, obj := range c.Objs {
		if !obj.Matches(filter) {
			continue
		}
		util.DebugLog(util.DebugLevelApi, "Show {{.Name}}", "obj", obj)
		err := cb(obj)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *{{.Name}}Cache) SetNotifyCb(fn func(obj *{{.Name}}Key)) {
	c.notifyCb = fn
}
{{- end}}

`

func (m *mex) generateMessage(file *generator.FileDescriptor, desc *generator.Descriptor) {
	message := desc.DescriptorProto
	if GetGenerateMatches(message) && message.Field != nil {
		m.P("func (m *", message.Name, ") Matches(filter *", message.Name, ") bool {")
		m.P("if filter == nil { return true }")
		for _, field := range message.Field {
			m.generateFieldMatches(message, field)
		}
		m.P("return true")
		m.P("}")
		m.P("")
	}
	if HasGrpcFields(message) {
		m.generateFields([]string{*message.Name + "Field"}, []string{}, desc)
		m.P("")
		m.P("var ", *message.Name, "AllFields = []string{")
		m.generateAllFields([]string{*message.Name + "Field"}, []string{}, desc)
		m.P("}")
		m.P("")
	}
	msgtyp := m.gen.TypeName(desc)
	m.P("func (m *", msgtyp, ") CopyInFields(src *", msgtyp, ") {")
	if HasGrpcFields(message) {
		m.P("fmap := make(map[string]struct{})")
		m.P("// add specified fields and parent fields")
		m.P("for _, set := range src.Fields {")
		m.P("for {")
		m.P("fmap[set] = struct{}{}")
		m.P("idx := strings.LastIndex(set, \".\")")
		m.P("if idx == -1 { break }")
		m.P("set = set[:idx]")
		m.P("}")
		m.P("}")
		m.importStrings = true
	}
	m.generateCopyIn(make([]string, 0), make([]string, 0), desc, make([]*generator.Descriptor, 0), HasGrpcFields(message))
	m.P("}")
	m.P("")

	if GetGenerateCud(message) {
		if !HasMessageKey(message) {
			m.gen.Fail("message", *message.Name, "needs a unique key field named key of type", *message.Name+"Key", "for option generate_cud")
		}
		args := cudTemplateArgs{
			Name:      *message.Name,
			CudName:   *message.Name + "Cud",
			KeyName:   *message.Name + "Key",
			HasFields: HasGrpcFields(message),
			GenCache:  GetGenerateCache(message),
		}
		m.cudTemplate.Execute(m.gen.Buffer, args)
		m.importUtil = true
	}
	if GetObjKey(message) {
		m.P("func (m *", message.Name, ") GetKeyString() string {")
		m.P("key, err := json.Marshal(m)")
		m.P("if err != nil {")
		m.P("util.FatalLog(\"Failed to marshal ", message.Name, " key string\", \"obj\", m)")
		m.P("}")
		m.P("return string(key)")
		m.P("}")
		m.P("")

		m.P("func ", message.Name, "StringParse(str string, key *", message.Name, ") {")
		m.P("err := json.Unmarshal([]byte(str), key)")
		m.P("if err != nil {")
		m.P("util.FatalLog(\"Failed to unmarshal ", message.Name, " key string\", \"str\", str)")
		m.P("}")
		m.P("}")
		m.P("")
		m.importUtil = true
	}
	if HasMessageKey(message) {
		m.P("func (m *", message.Name, ") GetKey() *", message.Name, "Key {")
		m.P("return &m.Key")
		m.P("}")
		m.P("")
	}
}

func (m *mex) generateService(file *generator.FileDescriptor, service *descriptor.ServiceDescriptorProto) {
	if len(service.Method) != 0 {
		for _, method := range service.Method {
			m.generateMethod(file, service, method)
		}
	}
}

func (m *mex) generateMethod(file *generator.FileDescriptor, service *descriptor.ServiceDescriptorProto, method *descriptor.MethodDescriptorProto) {

}

func HasGrpcFields(message *descriptor.DescriptorProto) bool {
	if message.Field != nil && len(message.Field) > 0 && *message.Field[0].Name == "fields" && *message.Field[0].Type == descriptor.FieldDescriptorProto_TYPE_STRING {
		return true
	}
	return false
}

func HasMessageKey(message *descriptor.DescriptorProto) bool {
	if message.Field == nil {
		return false
	}
	if len(message.Field) > 0 && *message.Field[0].Name == "key" {
		return true
	}
	if len(message.Field) > 1 && HasGrpcFields(message) && *message.Field[1].Name == "key" {
		return true
	}
	return false
}

func GetGenerateMatches(message *descriptor.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, protogen.E_GenerateMatches, false)
}

func GetGenerateCud(message *descriptor.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, protogen.E_GenerateCud, false)
}

func GetGenerateCache(message *descriptor.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, protogen.E_GenerateCache, false)
}

func GetObjKey(message *descriptor.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, protogen.E_ObjKey, false)
}
