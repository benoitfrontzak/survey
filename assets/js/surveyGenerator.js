const UI = new SurveyGenerator_UI(),
      API = new SurveyGenerator_API()



window.addEventListener('DOMContentLoaded', (event) => { 
    let totalQuestions = 1,
        totalRules = 1
// NAVIGATION 
    // Next
    $('#firstNext').bind('click', () =>{
        generalInformationIsValid = UI.validGeneralInformation()
        if(generalInformationIsValid == true){
            UI.showNav('second')
        }        
    })
    $('#secondNext').bind('click', () =>{
        questionnaireIsValid = UI.validQuestionnaire()
        if(questionnaireIsValid == true){
            UI.showNav('third')
        }        
    })

    // Previous
    $('#firstPrevious').bind('click', () =>{ UI.showNav('first') })
    $('#secondPrevious').bind('click', () =>{ UI.showNav('second') })

// QUESTIONNAIRE
    // When answers number change
    $('#answersNumber').on('change', () => {
        const answersNumber = $('#answersNumber').val()
        UI.answersNumberIsValid(answersNumber)
    })

    // When add checkbox button is clicked
    $('#addCheckbox').bind('click', () =>{ 
        const answersNumber = $('#answersNumber').val() * 1
        const valid = UI.answersNumberIsValid(answersNumber)
        if(valid){
            UI.addCheckboxAnswers(totalQuestions, answersNumber)
            document.querySelector('#totalQuestions').value = totalQuestions
            totalQuestions += 1
        }
        
    })

    // When add radio button is clicked
    $('#addRadio').bind('click', () =>{ 
        const answersNumber = $('#answersNumber').val()
        const valid = UI.answersNumberIsValid(answersNumber)
        if(valid){
            UI.addRadioAnswers(totalQuestions, answersNumber)
            document.querySelector('#totalQuestions').value = totalQuestions
            totalQuestions += 1
        }
    })

    // When add input button is clicked
    $('#addInput').bind('click', () =>{ 
        const answersNumber = $('#answersNumber').val()
        const valid = UI.answersNumberIsValid(answersNumber)
        if(valid){
            UI.addFreeAnswer(totalQuestions, answersNumber)
            document.querySelector('#totalQuestions').value = totalQuestions
            totalQuestions += 1
        }
    })

// REPORT
    // When operator lower than is clicked
    $('#addLower').bind('click', () =>{ 
        const ruleValue = $('#ruleValue').val() * 1
        const valid = UI.ruleValueIsValid(ruleValue)
        if(valid){
            UI.addOperatorRule(totalRules, ruleValue, 'lower than')
            document.querySelector('#totalRules').value = totalRules
            totalRules += 1
        }
        
    })
    // When operator equal to is clicked
    $('#addEqual').bind('click', () =>{ 
        const ruleValue = $('#ruleValue').val() * 1
        const valid = UI.ruleValueIsValid(ruleValue)
        if(valid){
            UI.addOperatorRule(totalRules, ruleValue, 'equal to')
            document.querySelector('#totalRules').value = totalRules
            totalRules += 1
        }
        
    })
    // When operator greater than is clicked
    $('#addGreater').bind('click', () =>{ 
        const ruleValue = $('#ruleValue').val() * 1
        const valid = UI.ruleValueIsValid(ruleValue)
        if(valid){
            UI.addOperatorRule(totalRules, ruleValue, 'greater than')
            document.querySelector('#totalRules').value = totalRules
            totalRules += 1
        }
        
    })

    // When save button (submit) is clicked
    $('#save').bind('click', (e) => {
        e.preventDefault()
        const form = document.querySelector('#surveyGeneratorForm'),
              totalRules = $('#totalRules').val(),
              reportIsValid = UI.validateReport(totalRules)
        if (reportIsValid == true){
            form.submit()
        }
        // let data = new FormData(form);
        // let formObj = UI.serialize(data);
        // console.log("formObj");
        // console.log(formObj);
        // API.addNewSurvey(formObj).then(resp => {
        //     console.log('resp')
        //     console.log(resp)
        // })
    })
})