module backend

go 1.23

toolchain go1.23.0

require (
	github.com/HansRobo/ssl_gui_test v0.0.0-20240828231819-9ad5015d0cd0
	// requireセクションは、プロジェクトが依存するパッケージを定義します。
	github.com/gorilla/websocket v1.5.3
)

require google.golang.org/protobuf v1.34.2 // indirect
