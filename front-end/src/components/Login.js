import { useState } from "react";
import Input from "./form/Input.js"
import { useNavigate, useOutletContext } from "react-router-dom";
const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const { setJwtToken } = useOutletContext();
    const { setAlertClassname } = useOutletContext();
    const { setAlertMessage } = useOutletContext();
    const navigate = useNavigate();
    const handleSubmit = (evnet) => {
        evnet.preventDefault();
        console.log("email/pass", email, password);
        if (email == "admin@example.com") {
            setJwtToken("abc");
            setAlertClassname("d-none");
            setAlertMessage("")
            navigate("/")
        } else{
            
            setAlertClassname("alert-danger");
            setAlertMessage("Invalid credentials")
        }

    }

    return (
        <div className="col-md-6 offset-md-3">
            <h2>Login</h2>
            <hr />

            <form onSubmit={handleSubmit}>
                <Input
                    title="Email Address"
                    type="email"
                    className="form-control"
                    name="email"
                    autoComplete="email-new"
                    onChange={(event) => setEmail(event.target.value)}
                />

                <Input
                    title="Password"
                    type="password"
                    className="form-control"
                    name="password"
                    autoComplete="password-new"
                    onChange={(event) => setPassword(event.target.value)}
                />
                <input
                    type="submit"
                    className="btn btn-primary"
                    value="Login"
                />
            </form>
        </div>
    )
}

export default Login;