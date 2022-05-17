import { memo, ReactNode } from "react"
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

type Props = {
    onClose: () => void;
    isOpen: boolean;
}


export const SuggestModal = memo((props: Props) => {

    const { onClose, isOpen } = props;

    const onClickSuggest = () => {
        console.log("test")
    }

    return(
        <Modal
            isOpen={isOpen}
            onClose={onClose}
        >
            <ModalOverlay />
            <ModalContent>
            <ModalHeader>Suggestion</ModalHeader>
            <ModalCloseButton />
            <ModalBody pb={6}>
                {/* <FormControl>
                <FormLabel>First name</FormLabel>
                <Input placeholder='First name' />
                </FormControl>

                <FormControl mt={4}>
                <FormLabel>Last name</FormLabel>
                <Input placeholder='Last name' />
                </FormControl> */}
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