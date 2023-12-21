package main

func main() {
    router := setupRouter()
    router.Run(":8080")
}
