import React from "react";
import List from "./list";

export default function MainPage() {
    return (
        <div className="container">
            <p className="title">
                Data received from backend
            </p>

            <List/>
        </div>
    )
}