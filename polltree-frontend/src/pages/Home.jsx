import { IoLogInOutline } from "react-icons/io5";
import { VscSend } from "react-icons/vsc";

import "./Home.css"

export default function Home() {
    return (
        <div className="home_main">
            <div className="home_hero">
                <div className="hero_l1">
                    <p>
                        v0.0.1
                    </p>
                </div>
                <div className="hero_h1">
                    <p>Know what people care about</p>
                </div>
                <div className="hero_h2">
                    <p>
                        get to know what people think but giving a poll with them and let them reply
                    </p>
                </div>

                <div className="hero_b2">
                    <a>
                        Make a Poll
                        <VscSend />
                    </a>
                    <a>
                        Fund
                        <IoLogInOutline />
                    </a>
                </div>
            </div>
        </div>
    )
}