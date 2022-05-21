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
    Heading,
    Box,
  } from '@chakra-ui/react'
import { useMessage } from "../../hooks/useMessage";
import axios from "axios";
import { useHistory } from "react-router-dom";
import { useLoginUser } from "../../hooks/useLoginUser";
import DatePicker from "react-datepicker"
import { Place } from "../../types/place"
import { Map } from "../molecules/Map"

import "react-datepicker/dist/react-datepicker.css"

type Props = {
    onClose: () => void;
    isOpen: boolean;
    address: string;
    name: string;
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

    const [startDate, setStartDate] = useState(new Date());

    const { onClose, isOpen, address, name } = props;

    const onClickPost = () => {
        console.log(startDate)

        var date = startDate.getFullYear() + "/" + ("00" + (startDate.getMonth()+1)).slice(-2) + "/" + ("00" + startDate.getDate()).slice(-2);
        console.log(date)

        var params = new URLSearchParams();
        params.append('place_id', '3');
        params.append('point', point);
        params.append('uesrname', loginUser!);
        params.append('date', date);
        axios.post("http://localhost:8080/post/new",params)
        .then((res) => {
            if(res.status == 200){
                console.log("post success")
                showMessage({title: "投稿完了", status:"success"})
                history.push("/home")
            }
        }).catch(() => {
            console.log("post failed")
            showMessage({title:"投稿失敗", status:"error"})
        })
    }

    return(
        <Modal
            isOpen={isOpen}
            onClose={onClose}
        >
            <ModalOverlay />
            <ModalContent>
            <ModalHeader>新規投稿</ModalHeader>
            <ModalCloseButton />
            <ModalBody >
                <FormControl>
                    <Heading size="sm">{name}</Heading>

                    <Heading size="sm" pt={5} pb={2}>日付</Heading>
                    <DatePicker selected={startDate} dateFormat="yyyy/MM/dd" onChange={(date:Date) => setStartDate(date)} />

                    <Heading size="sm" pt={5} pb={2}>ポイント</Heading>
                    <Input value={point} onChange={onChangePoint} placeholder='point'/>

                    <Box pt={3}>
                        <Map place={address}/>
                    </Box>
                    

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