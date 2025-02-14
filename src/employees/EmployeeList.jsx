import { useEffect, useState } from "react";
import PageHeader from "../header/PageHeader";
import axios from 'axios';

function EmployeeList() {
    const [employees, setEmployees] = useState([]);
    const readAllEmployees = async () => {
        try {
            const baseUrl = 'http://localhost:8081'
            const response = await axios.get(`${baseUrl}/employees`);
            setEmployees(response.data);
        } catch(error) {
            alert('Server Error');
        }
    };
    useEffect(()=>{ readAllEmployees(); },[]);
    const OnDelete = async (id) => {
        if(!confirm('Are you sure?')) return;
        try {
            const baseUrl = 'http://localhost:8081'
            const response = await axios.delete(`${baseUrl}/employees/${id}`);
            alert(response.data.message)
            readAllEmployees();
        } catch(error) {
            alert('Server Error');
        }
    }
    return (
        <>
            <PageHeader PageNumber={1}/>
            <h3>List of Employees</h3>
            <div className="container">
                <table className="table table-primary table-striped">
                    <thead className="table-dark">
                        <tr>
                            <th scope="col">Employee Name</th>
                            <th scope="col">Position</th>
                            <th scope="col">Department</th>
                            <th scope="col">Salary</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        { employees.map( (employee) => {
                            return (
                            <tr>
                                <th scope="row">{employee.name}</th>
                                <td>{employee.position}</td>
                                <td>{employee.department}</td>
                                <td>{employee.salary}</td>
                                <td>
                                    <a href={`/employees/edit/${employee.id}`} className="btn btn-warning">Edit Salary</a>
                                    <button className="btn btn-danger"
                                        onClick={()=>{OnDelete(employee.id);}}>Delete</button>
                                </td>
                            </tr>
                            );
                        } ) 
                        }
                    </tbody>
                </table>
            </div>
        </>
    );
}

export default EmployeeList;
