import {memo, VFC} from "react"
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

import { PostModal } from "../organisms/PostModal";

export const PostNew: VFC = memo(() => {

    const history = useHistory()
    const { loginUser } = useLoginUser();

    const { isOpen, onOpen, onClose } = useDisclosure()

    if (loginUser === null){
        history.push("/login");
    }

    return(
        <>
            <p>New Post</p>
            <p>お店を検索する</p>

            <Button colorScheme='teal' onClick={onOpen} autoFocus={false}>post modal</Button>
            <PostModal onClose={onClose} isOpen={isOpen} />
        </>
    )
});