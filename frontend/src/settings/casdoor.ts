import Sdk from "casdoor-js-sdk";
import Conf from '../config/casdoor.ts'

const setToken = (token:string) => {
    localStorage.setItem("token", token)
}
const removeToken = (token: string) => {
    localStorage.removeItem(token)
}

const Setting = {
    sdk: new Sdk(Conf),
    setToken,
    removeToken,
}


export default Setting
