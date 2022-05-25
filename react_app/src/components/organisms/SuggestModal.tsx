import { memo, ReactNode, useState, useEffect } from "react"
import { Modal, ModalOverlay, ModalContent, ModalHeader, ModalFooter, ModalBody, ModalCloseButton, FormControl, FormLabel, Input, Button, Heading } from '@chakra-ui/react'

import { Map } from "../molecules/Map"
import axios from "axios";
import { useLoginUser } from "../../hooks/useLoginUser";
import { Suggest } from "../../types/suggest";

type Props = {
    onClose: () => void;
    isOpen: boolean;
}


export const SuggestModal = memo((props: Props) => {

    const { onClose, isOpen } = props;
    const [ address, setAddress ] = useState<string>("");
    const [ place, setPlace ] = useState<Suggest>();
    const { loginUser } = useLoginUser();



    const onClickSuggest = () => {
        setAddress("茨城県つくば市天久保２丁目６−１")
        axios.get("http://localhost:8080/place/suggest", {params: {username: loginUser}})
        .then((res) => {
            if(res.status === 200){
                console.log("success")
                setPlace(res.data)
                console.log(res.data)
            }
        }).catch(() => {
            console.log("error")
        })
    }

    useEffect(() => {
        onClickSuggest()
        console.log("test")
    },[isOpen])

    return(
        <Modal isOpen={isOpen} onClose={onClose} size={'3xl'}>
            <ModalOverlay />
            <ModalContent>
            <ModalHeader>このお店はどうでしょうか？</ModalHeader>
            <ModalCloseButton />
            <ModalBody pb={6}>
                { place ?
                    <>
                    <Heading>{place!.place.name}</Heading>
                    <Heading>{place!.value}</Heading>
                    <Map address={place!.place.address} name={place!.place.name} />
                    </>
                : <></>}
            </ModalBody>

            <ModalFooter>
                <Button colorScheme='blue' mr={3} onClick={onClickSuggest}>
                次へ
                </Button>
                <Button onClick={onClose}>閉じる</Button>
            </ModalFooter>
            </ModalContent>
        </Modal>
    )
})