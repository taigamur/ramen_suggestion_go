import { useCallback, useState } from "react"
import axios from "axios"
import {Place} from "../types/place"

export const getAllPlaces = () => {
    // const { showMessage } = useMessage()
    // const [loading, setLoading] = useState(false)

    const [places, setPlaces] = useState<Array<Place>>([])

    const getPlaces = useCallback(() => {
        // axios.get<Array<Place>>("https://jsonplaceholder.typicode.com/users")
        axios.get<Array<Place>>("http://localhost:8080/places/index")
        .then((res) => {
            setPlaces(res.data)
            // showMessage({title: "ユーザーの取得に成功", status: "error"})
        })
        .catch(() => {
            // showMessage({title: "ユーザーの取得に失敗", status: "error"})
        })
    }, [])
    return { getPlaces, places}
}