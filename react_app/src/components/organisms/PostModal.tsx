import { memo, useState, ChangeEvent } from "react"
import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalFooter,
    ModalBody,
    ModalCloseButton,
    FormControl,
    FormLabel,
    Input,
    Button,
  } from '@chakra-ui/react'
import { useMessage } from "../../hooks/useMessage";
import axios from "axios";
import { useHistory } from "react-router-dom";
import { useLoginUser } from "../../hooks/useLoginUser";

type Props = {
    onClose: () => void;
    isOpen: boolean;
}


export const PostModal = memo((props: Props) => {

    const history = useHistory();
    const { showMessage } = useMessage();

    const { loginUser } = useLoginUser();

    if (loginUser === null){
        history.push("/login");
    }

    const [point, setPoint] = useState("");
    const onChangePoint = (e: ChangeEvent<HTMLInputElement>) => setPoint(e.target.value);

    const { onClose, isOpen } = props;

    const onClickPost = () => {
        console.log("post")
        var params = new URLSearchParams();
        params.append('place_id', '3');
        params.append('point', point);
        params.append('uesr_id', String(loginUser!.id))
        axios.post("http://localhost:8080/post/new",params)
        .then((res) => {
            if(res.status == 200){
                console.log("post success")
            }
        }).catch(() => {
            console.log("post failed")
        })
    }

    return(
        <Modal
            isOpen={isOpen}
            onClose={onClose}
        >
            <ModalOverlay />
            <ModalContent>
            <ModalHeader>New Post</ModalHeader>
            <ModalCloseButton />
            <ModalBody pb={6}>
                <FormControl>

                    <FormLabel>yyyy/mm/dd</FormLabel>
                    <Input id="date" placeholder='date' />

                    <FormLabel>Place</FormLabel>
                    <Input readOnly placeholder='****'/>

                    <FormLabel>Point</FormLabel>
                    <Input value={point} onChange={onChangePoint} placeholder='point' />

                </FormControl>
            </ModalBody>

            <ModalFooter>
                <Button colorScheme='blue' mr={3} onClick={onClickPost}>
                Post
                </Button>
                <Button onClick={onClose}>Cancel</Button>
            </ModalFooter>
            </ModalContent>
        </Modal>
    )
})