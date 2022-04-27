class App_API {

  // GET country information of given country name
  // We are interested in getting Language(Name, Code)
  // and Currency(Name, Code, Symbol)
  async countryInfo(countryName){
    const url = 'https://restcountries.com/v3.1/name/' + countryName
    const response = await fetch(url)
    const data = await response.json()
    return data
  }
  // GET companyname is uniq
  async companynameIsUniq(companyname){
    const url = "/companyIsUniq/" + companyname
    const response = await fetch(url)
    const data = await response.json()
    return data
  }
  // GET placename is uniq
  async placenameIsUniq(placename){
    const url = "/placeIsUniq/" + placename
    const response = await fetch(url)
    const data = await response.json()
    return data
  }
  // GET autoclean decrypted tmp files
  async autocleanDecryptedFiles(activeHash){
    const url = "/autoclean/" + activeHash
    const response = await fetch(url)
    const data = await response.json()
    return data
  }
  // POST new employee
  async addNewEmployee(serializedForm){
    const url = '/employee/add';
    const response = await fetch(url, {
      method: 'POST',
      body: JSON.stringify(serializedForm)
    })
    const result = await response.json()
    return result
  }
}