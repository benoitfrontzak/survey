const UI = new Survey_UI()

window.addEventListener('DOMContentLoaded', (event) => { 
    console.log('ok');
     // When save button (submit) is clicked
     $('#save').bind('click', (e) => {
        e.preventDefault()
        // Define the classname of fields to be validated
        // (validate = non empty or at least one checked)
        classnameList = ['checkboxes', 'radios', 'inputs']
        const form = document.querySelector('#surveyForm'),
              formIsValid = UI.validateSurveyForm(classnameList)    
        if(formIsValid == true){
            form.submit()
        }
 
    })
})