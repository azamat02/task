import React, {useEffect, useState} from "react";
import {testRequest} from "./services/main";
import {Simulate} from "react-dom/test-utils";
import input = Simulate.input;

export default function List() {
    const [labels, setLabels] = useState()
    const [errorMsg, setErrorMsg] = useState<String>()
    let result = []

    useEffect(()=>{
        testRequest().then((res)=>{
            console.log(res)
            setLabels(res.Labels)
        }).catch((err)=>{
            console.log(err)
            setErrorMsg('Произошла ошибка, попробуйте позже')
        })
    }, [setLabels])

    if (!labels) {
        return <>Loading</>
    }

    if (errorMsg) {
        return <>{errorMsg}</>
    }

    for (const [key, value] of Object.entries(labels!)) {
        result.push({label: key, value: value})
    }

    const renderedList = result.map((label, index) => {
        return (
            <div className={"list-item"} key={`${index}-label`}>
                <p className="label">
                    {label.label}
                </p>
                <input type="text" className="input" defaultValue={''+label.value}/>
            </div>
        )
    })

    return (
        <div className="inputs">
            {renderedList}
        </div>
    )
}