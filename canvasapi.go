package main


const BASE string = "https://canvas.instructure.com/api/v1/"

type CanvasApi struct {
    auth string
}


func NewCanvasApi(auth string) *CanvasApi {
    return &CanvasApi {
        auth, 
    }
}

