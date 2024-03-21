import Setting from "../settings/casdoor.ts";
import sdkConfig from "../config/casdoor.ts";

const LoginPage = () => {
    const login = (
        organizationName: string = sdkConfig.organizationName,
        client_id: string = sdkConfig.clientId,
        response_type: string | number,
        redirect_uri: string = sdkConfig.redirectPath,
        scope: string = 'read',
        state: string,
    ) => {
        // app: app-built-in默认
        // client_id=014ae4bd048734ca2dea&response_type=code&redirect_uri=https%3A%2F%2Fforum.casbin.com%2Fcallback&scope=read&state=app-casnode
        //fetch(`http://192.168.2.185:8000/signup/?or,ganizationName=${organizationName}&client_id=${client_id}&response_type=${response_type}&redirect_uri=${redirect_uri}&scope=${scope}&state=${state}`, {
        fetch(`http:/localhost:9000/callback`,{
            method: "GET",
        }).then((res) => {
            console.log(res.body)
        }).catch(err => {
            console.error(err)
        })
    }
    return <>
        <button
            onClick={() => Setting.sdk.getSigninUrl()}
        >Login
        </button>

        <button
            onClick={() => login}
        >Login
        </button>
    </>
}

export default LoginPage
