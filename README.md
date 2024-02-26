一、工程目录结构
/
├── api
├── hack
├── internal
│   ├── cmd
│   ├── consts
│   ├── controller
│   ├── dao
│   ├── logic
│   ├── model
│   |   ├── do
│   │   └── entity
│   └── service
├── manifest
├── resource
├── utility
├── go.mod
└── main.go 

🔥重要提示🔥：框架的工程目录采用了通用化的设计，以满足不同复杂程度业务项目的需求，但实际项目中可以根据项目需要适当增减默认的目录。例如，没有i1
8n/template/protobuf需求的场景，直接删除对应目录即可。又例如，非常简单的业务项目（如验证/演示项目），不考虑使用严谨的dao/logic/model目录
及特性，那么直接删除对应目录即可，可以在controller中直接实现业务逻辑。一切都可以由开发者灵活选择组装！

关于工程目录结构说明详见gf官方文档：https://goframe.org/pages/viewpage.action?pageId=30740166