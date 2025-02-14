import { useState } from "react";
import PageHeader from "../header/PageHeader";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function EmployeeCreate() {
    const [employee, setEmployee] = useState({id:'', name:'', position:'',
            department:'', salary: 0.0, email: '', phone: ''});
    const OnBoxChange = (event) => {
        const newEmployee = {...employee};
        newEmployee[event.target.id] = event.target.value;
        setEmployee(newEmployee);
    }
    const navigate = useNavigate();
    const OnCreate = async () => {
        try {
            const baseUrl = 'http://localhost:8081'
            const response = await axios.post(`${baseUrl}/employees`, {...employee,
                                        salary:parseFloat(employee.salary)});
            alert(response.data.message)
            navigate('/employees/list');
        } catch(error) {
            alert('Server Error');
        }
    }
    return (
        <>
            <PageHeader PageNumber={2}/>
            <h3><a href="/employees/list" className="btn btn-light">Go Back</a>New Employee</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="name" className="form-label">Employee Name:</label>
                    <input type="text" className="form-control" id="name" 
                        placeholder="Please enter employee name"
                        value={employee.name} onChange={OnBoxChange}/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="position" className="form-label">Position:</label>
                    <input type="text" className="form-control" id="position" 
                        placeholder="Please enter position"
                        value={employee.position} onChange={OnBoxChange}/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="department" className="form-label">Department:</label>
                    <input type="text" className="form-control" id="department" 
                        placeholder="Please enter department"
                        value={employee.department} onChange={OnBoxChange}/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="salary" className="form-label">Salary:</label>
                    <input type="text" className="form-control" id="salary" 
                        placeholder="Please enter salary"                        
                        value={employee.salary} onChange={OnBoxChange}/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="email" className="form-label">Email:</label>
                    <input type="text" className="form-control" id="email" 
                        placeholder="Please enter email"                                                 
                        value={employee.email} onChange={OnBoxChange}/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="phone" className="form-label">Phone:</label>
                    <input type="text" className="form-control" id="phone" 
                        placeholder="Please enter phone"                                                 
                        value={employee.phone} onChange={OnBoxChange}/>
                </div>
                <button className="btn btn-success"
                    onClick={OnCreate}>Create Employee</button>
            </div>
        </>
    );
}

export default EmployeeCreate;
