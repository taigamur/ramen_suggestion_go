import { memo, VFC } from "react"
import { Box, Stack, Text } from "@chakra-ui/react"

import { Post } from "../../types/post"

type Props = {
    post: Post;
}
export const PostItem: VFC<Props> = memo((props) => {
    const { post } = props;


    return(
        <Box w='100%' p='6'>
            <Text>{post.id}</Text>
            <Text>{post.username}</Text>
            <Text>{post.place_id}</Text>
            <Text>{post.place.name}</Text>
        </Box>
    )
})