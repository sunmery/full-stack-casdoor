import Sdk from "casdoor-js-sdk";
import Conf from '../config/casdoor.ts'

const setToken = (token:string) => {
    localStorage.setItem("token", token)
}
const removeToken = (token: string) => {
    localStorage.removeItem(token)
}

const isLoggedIn = () => {
    const token = localStorage.getItem("token");
    return token !== null && token.length > 0;
};

const goToLink = (link:string) => {
    location.href = link;
};

const Setting = {
    sdk: new Sdk(Conf),
    setToken,
    removeToken,
    isLoggedIn,
    goToLink,
}



export default Setting
