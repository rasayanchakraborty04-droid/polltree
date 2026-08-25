import { FaGithub } from "react-icons/fa";
import { FaXTwitter } from "react-icons/fa6";

import "./Navbar.css"

export default function Navbar() {
    return (
        <navbar>
            <div className="navbar_head">
                <p>
                    poll<span>tree</span>
                </p>
            </div>

            <div className="navbar_options">
                <p>tutorial</p>
                <p>pricing</p>
                <a href="/login">signup</a>
            </div>

            <div className="navbar_icons">
                <FaXTwitter />
                <FaGithub />
            </div>
        </navbar>
    )
}