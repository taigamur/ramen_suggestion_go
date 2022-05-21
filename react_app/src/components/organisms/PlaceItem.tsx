import { memo, VFC } from "react"
import { Box, Stack, Text } from "@chakra-ui/react"

import { Place } from "../../types/place"

type Props = {
    place: Place;
}
export const PostItem: VFC<Props> = memo((props) => {
    const { place } = props;


    return(
        <Box w='100%'>
            <Text>{place.id}</Text>
            <Text>{place.name}</Text>
            <Text>{place.address}</Text>
        </Box>
    )
})