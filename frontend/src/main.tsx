
import ReactDOM from "react-dom/client";
import App from "./App";
import "./styles.css";
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import CallbackPage from "./pages/CallbackPage.tsx";

const router = createBrowserRouter([
    {
        path: "/",
        element: <App/>,
    },
    {
        path: "/callback",
        element: <CallbackPage/>
    }
]);
ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
    // <React.StrictMode>
    <RouterProvider router={router}/>

    // </React.StrictMode>,
);
