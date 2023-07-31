//? REACT
import React, { useEffect, useState } from 'react';

//? AXIOS
import axios from 'axios';

const Store = () => {
  const [htmlContent, setHtmlContent] = useState('');

  useEffect(() => {
    const getPartTwo = async () => {
      try {
        const response = await axios.get('http://localhost:8080/');
        setHtmlContent(response.data);
      } catch (e) {
        console.log("Error fetching html", e);
      }
    }

    getPartTwo();
  }, []);

  return (
    <>
      <div dangerouslySetInnerHTML={{ __html: htmlContent }}></div>
    </>
  );
}

export default Store;
