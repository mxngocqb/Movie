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
        // build the requet payload
        let payload = {
            email: email,
            password: password,
        }

        const requestOptions = {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload),
            credentials: 'include'
        }

        fetch(`http://localhost:8080/authenticate`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                if (data && data.error) {
                    setAlertClassname("alert-danger")
                    setAlertMessage(data.message)
                } else {
                    setJwtToken(data.access_token)
                    setAlertClassname("d-none")
                    setAlertMessage("")
                    navigate("/")
                }
            })
            .catch(error => {
                setAlertClassname("alert-danger")
                setAlertMessage(error)
            })
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