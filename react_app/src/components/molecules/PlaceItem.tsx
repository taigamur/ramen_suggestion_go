import { memo, VFC } from "react"
import { Box, Stack, Text, Button } from "@chakra-ui/react"

type Props = {
    id: number;
    name: string;
    address: string
    onClick: (id: number) => void;
}

export const PlaceItem: VFC<Props> = memo((props) => {
    const { id, name, address, onClick } = props;


    return(
        <Button colorScheme='green' variant='outline' width='80%' mx="auto" onClick={() => onClick(id)} >
            <Box w='100%'>
                <Text>{name}</Text>
            </Box>
        </Button>
    )
})