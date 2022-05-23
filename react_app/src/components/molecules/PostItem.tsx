import { memo, VFC } from "react"
import { Box, Stack, Text, Button, IconButton, Spacer, Flex } from "@chakra-ui/react"
import { DeleteIcon, EditIcon } from "@chakra-ui/icons"


import { Post } from "../../types/post"

type Props = {
    post: Post;
}
export const PostItem: VFC<Props> = memo((props) => {
    const { post } = props;


    return(
        <Box w="80%" mx="auto" borderWidth='3px' borderRadius='lg'>
            <Flex>
            <Box w="30%">
                <Text>date: {post.date}</Text>
            </Box>
            <Spacer/>
            <Box w="70%">
                <Text>user: {post.username}</Text>
                <Text>place: {post.place.id}</Text>
                <Text>place name: {post.place.name}</Text>
                <Text>point: {post.value}</Text>  
            </Box>
            </Flex>

            <Box align="center">
                <IconButton mr={2} size='xs' variant='outline' colorScheme='teal' aria-label='Call Sage' icon={<EditIcon />}/>
                <IconButton size='xs' variant='outline' colorScheme='teal' aria-label='Call Sage' icon={<DeleteIcon />} />
            </Box>
        </Box>
    )
})