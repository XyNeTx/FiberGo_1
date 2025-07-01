import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import axios from 'axios'
import './App.css'

function App() {
  const [count, setCount] = useState(0)
  const [data, setData] =  useState<Employee[] | null>(null)

  class Employee{
    constructor(EmployeeID: string, Name: string , Age: number, Email: string, Salary: number){
      this.EmployeeID = EmployeeID
      this.Name = Name
      this.Age = Age
      this.Email = Email
      this.Salary = Salary
    }
    EmployeeID: string
    Name: string 
    Age: number
    Email: string
    Salary: number
  }


  useEffect(() => {
    const fetchData = async () =>{
      try{
        const response = await axios.get('http://localhost:8080/Employees/')
        setData(response.data)
        console.log(response.data)
      }
      catch(error){
        console.log(error)
      }
    }
    fetchData()
  }, [])

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
        <div>
            <h4>Data From API:</h4>
            <table className="table table-striped">
              <thead>
                <tr>
                  <th>Employee ID</th>
                  <th>Name</th>
                  <th>Age</th>
                  <th>Email</th>
                  <th>Salary</th>
                </tr>
              </thead>
              <tbody>
                {data && data.map((item: Employee) => (
                  <tr key={item.EmployeeID}>
                    <td>{item.EmployeeID}</td>
                    <td>{item.Name}</td>
                    <td>{item.Age}</td>
                    <td>{item.Email}</td>
                    <td>{item.Salary}</td>
                  </tr>
                ))}
              </tbody>
            </table>
        </div>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
