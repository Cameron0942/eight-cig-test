import React, { useEffect, useState } from "react";
import axios from "axios";
import Box from "@mui/material/Box";
import { DataGrid } from "@mui/x-data-grid";
import { Button } from "@mui/material";

const Dashboard = () => {
  const [employeeInfo, setEmployeeInfo] = useState([]);

  useEffect(() => {
    // Function to make the Axios GET request
    const fetchData = async () => {
      try {
        const response = await axios.get("http://localhost:8080/api/dashboard");
        setEmployeeInfo(response.data);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData(); // Call the function to trigger the GET request when the component is loaded
  }, []);

  // Define the columns for the DataGrid
  const columns = [
    { field: "id", headerName: "ID", width: 70 },
    { field: "name", headerName: "Name", width: 150 },
    { field: "performance", headerName: "Performance", width: 130 },
    { field: "date", headerName: "Date", width: 600 },
  ];

  // Map employeeInfo data to the format expected by DataGrid
  const rows = employeeInfo.map((employee) => ({
    id: employee.ID,
    name: employee.Name,
    performance: employee.Performance,
    date: new Date(employee.Date),
  }));

  return (
    <>
      <h1 style={{ textAlign: "center" }}>Dashboard</h1>
      <p style={{ textAlign: "center" }}>
        Below are the results to a MySQL query with the given data from the test
        document.
      </p>
      <p style={{ margin: "0 auto", textAlign: "center", width: "100ch" }}>
        MySQL is queried using Axios to make an HTTP GET request to the GoLang
        backend server. The backend server then handles querying the MySQL
        database, and returns the data with the employee performance arranged
        from highest to lowest. I used a Material UI DataGrid to display the
        results below. By default the employee data is sorted by performance from highest to
        lowest, but can also be sorted by name and date.
      </p>
      <br />
      <Box style={{ width: "50%", margin: "0 auto" }}>
        <DataGrid
          rows={rows}
          columns={columns}
          pageSize={10}
          sx={{
            ".MuiDataGrid-columnHeader": {
              backgroundColor: "#f0f0f0",
              color: "#333",
            },
            ".MuiDataGrid-row": {
              backgroundColor: "#f7f7f7",
            },
          }}
        />
      </Box>
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          marginTop: 5,
        }}
      >
        <Button
          variant="contained"
          onClick={() => {
            window.location.href = "/";
          }}
          sx={{
            background: "linear-gradient(#c0060d, #a31f25)", // Specify the linear gradient here for the button
            color: "#fff", // Set the text color to white for better visibility on the gradient background
            "&:hover": {
              color: "#71ff79", // Set the text color to green on hover
            },
          }}
        >
          Back to homepage
        </Button>
      </Box>
    </>
  );
};

export default Dashboard;
