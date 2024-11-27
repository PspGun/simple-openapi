import {Provider} from './components/ui/provider'
import {Box, Theme} from "@chakra-ui/react";
import {AuthPage} from "./pages/AuthPage.tsx";

function App() {

  return (
    <Provider>
        <Theme appearance="light">
            <Box gap={10} flexDirection={"column"} minH={"100vh"} minW={"100vw"} alignItems={"center"} display={"flex"} justifyContent={"center"} >
            <AuthPage />
            </Box>
        </Theme>
    </Provider>
  )
}

export default App
