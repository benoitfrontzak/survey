class SurveyGenerator_API{
    // POST new employee
    async addNewSurvey(serializedForm){
        const url = '/survey/generator';
        const response = await fetch(url, {
        method: 'POST',
        body: JSON.stringify(serializedForm)
        })
        const result = await response.json()
        return result
    }
}