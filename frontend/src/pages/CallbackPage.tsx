import {AuthCallback} from "casdoor-react-sdk";
import Setting from "../settings/casdoor.ts";
import sdkConfig from "../config/casdoor.ts";

export default function CallbackPage  ()  {
    return <>
        <AuthCallback
            sdk={Setting.sdk}
            serverUrl={sdkConfig.serverUrl}
            saveTokenFromResponse={(res) => {
                console.log(res);
                // Setting.setToken(res.data);
                // Setting.goToLink("/");
                location.href = "/"
            }}
            isGetTokenSuccessful={(res) => res.statusText === "ok"}
        />
    </>
}
