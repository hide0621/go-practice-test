# go-practice-test
参考：
https://docs.google.com/presentation/d/1p_fyQ7uSw6FYvsfEbKuZIKgVMNaWQ95JNV1dcXqMhbU/edit#slide=id.g4fa745498a_0_271

以下のコマンドを実行するとHTMLでテストのカバレッジを簡単に確認することができる

go test -coverprofile=out

go tool cover -html=out -o out.html

go tool cover -html=out
