import {memo, VFC, ReactNode} from "react"
import { Header } from "../organisms/Header";

// 受け取るpropsの型を定義
type Props = {
    children: ReactNode;
}

export const HeaderLayout: VFC<Props> = memo((props) => {
    const { children } = props;
    return (
        <>
            <Header/>
            {children}
        </>
    )
});