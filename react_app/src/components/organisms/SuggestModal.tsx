import { memo, ReactNode, useState, useEffect } from "react"
import { Modal, ModalOverlay, ModalContent, ModalHeader, ModalFooter, ModalBody, ModalCloseButton, FormControl, FormLabel, Input, Button } from '@chakra-ui/react'

import { Map } from "../molecules/Map"

type Props = {
    onClose: () => void;
    isOpen: boolean;
}


export const SuggestModal = memo((props: Props) => {

    const { onClose, isOpen } = props;
    const [ place, setPlace ] = useState<string>("");

    useEffect(() => {
        setPlace("茨城県つくば市天久保２丁目１１−１０")
    },[])

    const onClickSuggest = () => {
        console.log("test")
        setPlace("茨城県つくば市天久保２丁目６−１")
    }

    return(
        <Modal isOpen={isOpen} onClose={onClose} size={'3xl'}>
            <ModalOverlay />
            <ModalContent>
            <ModalHeader>Suggestion</ModalHeader>
            <ModalCloseButton />
            <ModalBody pb={6}>
                <Map place={place}/>
            </ModalBody>

            <ModalFooter>
                <Button colorScheme='blue' mr={3} onClick={onClickSuggest}>
                Next
                </Button>
                {/* <Button onClick={onClose}>Cancel</Button> */}
            </ModalFooter>
            </ModalContent>
        </Modal>
    )
})