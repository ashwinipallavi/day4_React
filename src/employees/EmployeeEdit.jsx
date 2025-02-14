import { useEffect, useState } from "react";
import PageHeader from "../header/PageHeader";
import { useNavigate, useParams } from "react-router-dom";
import axios from "axios";

function EmployeeEdit() {
    const [employee, setEmployee] = useState({id:'', name:'', position:'',
                department:'', salary: 0.0, email: '', phone: ''});
    const OnBoxChange = (event) => {
        const newEmployee = {...employee};
        newEmployee[event.target.id] = event.target.value;
        setEmployee(newEmployee);
    }
    const params = useParams();
    const readEmployeeById = async () => {
        alert(params.id);
        try {
            const baseUrl = 'http://localhost:8081'
            const response = await axios.get(`${baseUrl}/employees/${params.id}`);
            setEmployee(response.data);
        } catch(error) {
            alert('Server Error');
        }
    };
    useEffect(()=>{ readEmployeeById(); },[]);
    const navigate = useNavigate();
    const OnUpdate = async () => {
        try {
            const baseUrl = 'http://localhost:8081'
            const response = await axios.put(`${baseUrl}/employees/${params.id}`, {...employee,
                                        salary:parseFloat(employee.salary)});
            alert(response.data.message)
            navigate('/employees/list');
        } catch(error) {
            alert('Server Error');
        }
    }
    return (
        <>
            <PageHeader  PageNumber={1}/>
            <h3><a href="/employees/list" className="btn btn-light">Go Back</a>Edit Employee Salary</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="name" className="form-label">Employee Name:</label>
                    <div className="form-control" id="name">{employee.name}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="position" className="form-label">Position:</label>
                    <div className="form-control" id="position">{employee.position}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="department" className="form-label">Department:</label>
                    <div className="form-control" id="department">{employee.department}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="salary" className="form-label">Salary:</label>
                    <input type="text" className="form-control" id="salary" 
                        placeholder="Please enter salary"                                                 
                        value={employee.salary} onChange={OnBoxChange} />
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="email" className="form-label">Email:</label>
                    <div className="form-control" id="email">{employee.email}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="phone" className="form-label">Phone:</label>
                    <div className="form-control" id="phone">{employee.phone}</div>
                </div>
                <button className="btn btn-warning"
                    onClick={OnUpdate}>Update Salary</button>
            </div>
        </>
    );
}

export default EmployeeEdit;
