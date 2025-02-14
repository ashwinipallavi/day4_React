import { useState } from "react";
function FullName() {
       const [person, setPerson] = useState({first_name:'rahul', last_name:'dravid'});
       const onChangeFirstName = (event) => {
            const newPerson = {...person};
            newPerson.first_name = event.target.value;
            setPerson(newPerson)
       } 
       const onChangeBox = (event) => {
            const newPerson = {...person};
            newPerson[event.target.id] = event.target.value;
            setPerson(newPerson)
       } 
       return (
        <>
            <div className="container">
                <div className="form-group">
                    <label for="first_name">First Name:</label>
                    <input type="text" id="first_name"
                        className="form-control" value={person.first_name}
                        onChange={onChangeFirstName}/>
                </div>
                <div className="form-group">
                    <label for="last_name">Last Name:</label>
                    <input type="text" id="last_name"
                        className="form-control" value={person.last_name}
                        onChange={(event)=>setPerson({...person,last_name:event.target.value})}/>
                </div>
                <div>Full Name: {person.first_name} {person.last_name}</div>
            </div>
        </>
    )
}
export default FullName;