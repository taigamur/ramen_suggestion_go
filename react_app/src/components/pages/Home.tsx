import {memo, VFC, useCallback, useEffect, useState } from "react"
import { useHistory } from "react-router-dom";
import { useDisclosure, Button, Wrap, WrapItem } from '@chakra-ui/react'
import axios from "axios";

import { useLoginUser } from "../../hooks/useLoginUser";
import { SuggestModal } from "../organisms/SuggestModal";
import { Post } from "../../types/post"
import { PostItem } from "../molecules/PostItem"


export const Home: VFC = memo(() => {

    const history = useHistory()
    const { loginUser } = useLoginUser();

    const { isOpen, onOpen, onClose } = useDisclosure()
    const [ posts, setPosts ] = useState<Array<Post>>([])

    const onClickNewPost = useCallback(() => history.push("/post/new"),[]);

    const getPosts = () => {
        axios.get<Array<Post>>("http://localhost:8080/post/index", {params: {username: loginUser}})
        .then((res) => {
            console.log(res)
            setPosts(res.data)
        })
        .catch(() => {
            console.log("error")
        })
    }

    useEffect(() => getPosts(),[])

    return(
        <>
            <p>Homeページです。</p>
            <p>こんにちは、loginUser : {loginUser} さん</p>
    
            <Button colorScheme='teal' onClick={onOpen} autoFocus={false}>Suggestion</Button>
            <SuggestModal onClose={onClose} isOpen={isOpen}  />
            <Button colorScheme='teal' onClick={onClickNewPost} autoFocus={false}>NewPost</Button>

            <Wrap pt={10}>
                {posts.map((post) => (
                    <WrapItem key={post.id} w='100%' bg='green.200'>
                        <PostItem post={post}  />
                    </WrapItem>
                ))}
            </Wrap>
        </>
    )
});