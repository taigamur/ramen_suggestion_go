import {memo, VFC, useCallback, useEffect} from "react"
import { Redirect, useHistory } from "react-router-dom";
import { useLoginUser } from "../../hooks/useLoginUser";

import { useDisclosure, Button } from '@chakra-ui/react'
import { useCookies } from "react-cookie";

import { SuggestModal } from "../organisms/SuggestModal";

export const Home: VFC = memo(() => {

    const history = useHistory()
    const { loginUser } = useLoginUser();

    const { isOpen, onOpen, onClose } = useDisclosure()

    const onClickNewPost = useCallback(() => history.push("/post/new"),[]);

    // if(!loginUser){
    //     history.push("/login")
    // }

    return(
        <>
            <p>Homeページです。</p>
            <p>こんにちは、loginUser : {loginUser} さん</p>
    
            <Button colorScheme='teal' onClick={onOpen} autoFocus={false}>Suggestion</Button>
            <SuggestModal onClose={onClose} isOpen={isOpen}  />

            <Button colorScheme='teal' onClick={onClickNewPost} autoFocus={false}>NewPost</Button>
        </>
    )
});