import axios from "axios";
import { memo, VFC, useState, useEffect } from "react"
import { Wrap, WrapItem, Box } from "@chakra-ui/react"

import { User } from "../../types/user"
import { UserCard } from "../organisms/UserCard";

export const UserIndex: VFC = memo(() => {

    const [ users, setUsers ] = useState<Array<User>>([])

    const getUsers = () => {
        axios.get<Array<User>>("http://localhost:8080/users/index")
        .then((res) => {
            console.log(res)
            setUsers(res.data)
        })
        .catch(() => {
            console.log("error")
        });
    }

    const onClickUser = (id:number) => {

    }

    useEffect(() => getUsers(), [])

    return (
        <>
            <Wrap>
                {users.map((user) => (
                    <WrapItem key={user.id} mx="auto">
                    <UserCard id={user.id} name={user.name} onClick={onClickUser}/>
                    </WrapItem>
                ))}
            </Wrap>
        </>
    )
});