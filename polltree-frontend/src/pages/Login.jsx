import { IoIosArrowForward } from "react-icons/io";

import "./Login.css"

export default function Login() {
    return (
        <div className="login_main">
            <div className="login_box">
                <p className="login_box_header">tell us who you are</p>

                <div className="login_name">
                    <div className="login_field" id="login_firstName">
                        <IoIosArrowForward />
                        <input type="text" placeholder="first name" />
                    </div>
                    <div className="login_field" id="login_lastName">
                        <IoIosArrowForward />
                        <input type="text" placeholder="last name" />
                    </div>
                </div>

                <div className="login_email">
                    <div className="login_field" id="login_email">
                        <IoIosArrowForward />
                        <input type="email" placeholder="email" />
                    </div>
                </div>

                <div className="login_button">
                    <a href="/">Sign In</a>
                    <a href="/">Sign Up</a>
                </div>
            </div>
        </div>
    )
}