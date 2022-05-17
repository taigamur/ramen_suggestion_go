import {memo, VFC, useCallback} from "react"
import { useHistory } from "react-router-dom";
import { useLoginUser } from "../../hooks/useLoginUser";

import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalFooter,
    ModalBody,
    ModalCloseButton,
    Button,
    useDisclosure,
    FormControl,
    FormLabel,
    Input,

  } from '@chakra-ui/react'
import { SuggestModal } from "../organisms/SuggestModal";

export const Home: VFC = memo(() => {

    const history = useHistory()
    const { loginUser } = useLoginUser();

    if (loginUser === null){
        history.push("/login");
    }

    const { isOpen, onOpen, onClose } = useDisclosure()

    const onClickNewPost = useCallback(() => history.push("/post/new"),[]);

    return(
        <>
            <p>Homeページです。</p>
            <p>こんにちは、{loginUser?.name}さん</p>

            {/* <p>{username}のPost</p> */}
            <p>こんにちは、{loginUser?.name}さん</p>
    
            <Button colorScheme='teal' onClick={onOpen} autoFocus={false}>Suggestion</Button>
            <SuggestModal onClose={onClose} isOpen={isOpen}  />

            <Button colorScheme='teal' onClick={onClickNewPost} autoFocus={false}>NewPost</Button>
        </>
    )
});