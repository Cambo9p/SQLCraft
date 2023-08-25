/* this file takes the json from the api one a button is hit and displays it as a table */
document.getElementById('fetchButton').addEventListener('click', async () => {
  try {
    const response = await fetch('api/db/get_users')
    const responseData = await response.json()

    const tableBody = document.getElementById('table-body')
    tableBody.innerHTML = ''

    responseData.rows.forEach(rowData => {
      const tr = document.createElement('tr')
      rowData.forEach(cellData => {
        const td = document.createElement('td')
        td.textContent = cellData
        tr.appendChild(td)
      })
      tableBody.appendChild(tr)
    })
  } catch (error) {
    console.error('Error fetching data:', error)
  }
})
