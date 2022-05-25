import {memo, VFC, useCallback, useEffect, useState } from "react"
import { useHistory } from "react-router-dom";
import { useDisclosure, Button, Wrap, WrapItem, Box } from '@chakra-ui/react'
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

    const getPosts = (user: string) => {
        console.log("user: " + loginUser)
        axios.get<Array<Post>>("http://localhost:8080/post/index", {params: {username: loginUser}})
        .then((res) => {
            console.log(res)
            setPosts(res.data)
        })
        .catch(() => {
            console.log("error")
        })
    }

    useEffect(() => getPosts(loginUser!),[loginUser])

    return(
        <>
            <p>loginUser : {loginUser} さん</p>

            <Box w="100%" align='center' pb={3}>
                <Button colorScheme='teal' onClick={onOpen} autoFocus={false} variant='outline'>提案を見る</Button>
                <SuggestModal onClose={onClose} isOpen={isOpen} />
            </Box>

            <Box w="100%" align='center'>
                <Button colorScheme='teal' onClick={onClickNewPost} autoFocus={false}>新規投稿</Button>
            </Box>

            {posts ?
            <Wrap pt={5}>
                {posts.map((post) => (
                    <WrapItem key={post.id} width='100%'>
                        <PostItem post={post}  />
                    </WrapItem>
                ))}
            </Wrap>
            : <p></p>}
        </>
    )
});