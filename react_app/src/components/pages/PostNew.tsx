import { memo, VFC, useState, ChangeEvent } from "react"
import { useHistory } from "react-router-dom";
import { useLoginUser } from "../../hooks/useLoginUser";

import { Button, useDisclosure, FormControl, FormLabel, Input } from '@chakra-ui/react'

import { PostModal } from "../organisms/PostModal";
import axios from "axios";
import { Place } from "../../types/place"
import { useMessage } from "../../hooks/useMessage";

export const PostNew: VFC = memo(() => {

    const history = useHistory()
    const { showMessage } = useMessage();

    const { loginUser } = useLoginUser();
    const [ keyword, setKeyword ] = useState("");
    const onChangeKeyword = (e: ChangeEvent<HTMLInputElement>) => setKeyword(e.target.value);
    const [ places, setPlaces ] = useState<Array<Place>>([]);

    const { isOpen, onOpen, onClose } = useDisclosure()

    if (loginUser === null){
        history.push("/login");
    }

    const onClickFindPlace = () => {
        console.log(keyword)
        if ( keyword == ""){
            showMessage({title: "お店の名前を入力してください", status:"error"})
        }else{
            var params = new URLSearchParams();
            params.append('keyword', keyword)
            axios.post("http://localhost:8080/place/index", params)
            .then((res) => {
                if(res.status == 200){
                    console.log("success")
                    console.log(res.data)
                    setPlaces(res.data)
                }else{
                    console.log("failure")
                }
            })
        }
    }

    return(
        <>
            <p>New Post</p>
            <p>お店を検索する</p>

            <FormControl>
                <FormLabel></FormLabel>
                <Input placeholder='キーワード' value={keyword} onChange={onChangeKeyword}/>
                <Button onClick={onClickFindPlace}>検索</Button>
            </FormControl>

            <Button colorScheme='teal' onClick={onOpen} autoFocus={false}>post modal</Button>
            <PostModal onClose={onClose} isOpen={isOpen} />
        </>
    )
});