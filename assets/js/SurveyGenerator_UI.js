class SurveyGenerator_UI{
    // Navigation
    showNav(wanted){
        switch (wanted) {
            case 'first':
                $('#firstSpan').addClass('active')
                $('#secondSpan').removeClass('active')
                $('#thirdSpan').removeClass('active')

                $('#firstSpanLabel').show()
                $('#secondSpanLabel').hide()
                $('#thirdSpanLabel').hide()

                $('#surveyGenerator_GI').addClass('show active')
                $('#surveyGenerator_Q').removeClass('show active')
                $('#surveyGenerator_R').removeClass('show active')
                break;
        
            case 'second':
                $('#firstSpan').removeClass('active')
                $('#secondSpan').addClass('active')
                $('#thirdSpan').removeClass('active')

                $('#firstSpanLabel').hide()
                $('#secondSpanLabel').show()
                $('#thirdSpanLabel').hide()

                $('#surveyGenerator_GI').removeClass('show active')
                $('#surveyGenerator_Q').addClass('show active')
                $('#surveyGenerator_R').removeClass('show active')
                break;
                    
            case 'third':
                $('#firstSpan').removeClass('active')
                $('#secondSpan').removeClass('active')
                $('#thirdSpan').addClass('active')

                $('#firstSpanLabel').hide()
                $('#secondSpanLabel').hide()
                $('#thirdSpanLabel').show()

                $('#surveyGenerator_GI').removeClass('show active')
                $('#surveyGenerator_Q').removeClass('show active')
                $('#surveyGenerator_R').addClass('show active')
                break;

        }
    }

    // validate general information
    answersNumberIsValid(n){
        if( n != "" && !isNaN(n) ){
            // valid
            $('#answersNumber').removeClass('border border-danger')
            $('#answersNumberLabel').removeClass('text-danger fw-bold')
            return true
         }else{
            //  non valid
             $('#answersNumber').addClass('border border-danger')
             $('#answersNumberLabel').addClass('text-danger fw-bold')
             return false
         }
    }
    validGeneralInformation(){
        if($('#surveyTitle').val() != ""){
            $('#surveyTitle').removeClass('border border-danger')
            $('#surveyTitleLabel').removeClass('text-danger fw-bold')
            return true
        }else{
            $('#surveyTitle').addClass('border border-danger')
            $('#surveyTitleLabel').addClass('text-danger fw-bold')
            return false
        }
    }

    // validate questionnaire
    validQuestionnaire(){
        const totalQuestions = $('#totalQuestions').val()*1

        for (let index = 0; index < totalQuestions; index++) {
            const qLabel           = $('#questionLabel'+(totalQuestions)).val(),
                  qAnswersNumber   = $('#answersNumber'+(totalQuestions)).val()
            // qStory           = $('#questionStory'+(totalQuestions+1)).val(),
            // qHelper          = $('#questionLabel'+(totalQuestions+1)).val(),
            // qType            = $('#questionType'+(totalQuestions+1)).val()

            // Validate question's label
            const qLabelIsValid = this.validateQuestionLabel(qLabel, totalQuestions)

            // Store all answer's label isValid
            let allAnswersLabelValid = []
            // Store all answer's score isValid
            let allAnswersScoreValid = []

            // Validate answer's label & score
            for (let i = 0; i < qAnswersNumber; i++) {
                const aLabel      = $('#q'+(totalQuestions)+'answer'+i).val(),
                      aScore      = $('#q'+(totalQuestions)+'score'+i).val()
                // aComment    = $('#q'+(totalQuestions+1)+'comment'+i).val()

                const aLabelIsValid = this.validateAnswerLabel(aLabel,totalQuestions,i),
                      aScoreIsValid = this.validateAnswerScore(aScore,totalQuestions,i)

                allAnswersLabelValid.push(aLabelIsValid)
                allAnswersScoreValid.push(aScoreIsValid)
            }

            // Validate page
            if ( qLabelIsValid == false || allAnswersLabelValid.includes(false) || allAnswersScoreValid.includes(false) ){
                return false
            }else{
                return true
            }
        }
    }
    validateQuestionLabel(label, totalQuestions){
        if(label == '' || typeof(label) == 'undefined'){
            $('#questionLabel'+totalQuestions).addClass('border border-danger')
            $('#questionLabelLabel'+totalQuestions).addClass('text-danger fw-bold')
            return false
        }else{
            $('#questionLabel'+totalQuestions).removeClass('border border-danger')
            $('#questionLabelLabel'+totalQuestions).removeClass('text-danger fw-bold')
            return true
        }
    }
    validateAnswerLabel(label, totalQuestions, i){
        if(label == '' || typeof(label) == 'undefined'){
            $('#q'+(totalQuestions)+'answer'+i).addClass('border border-danger')
            $('#aLabelTitle').addClass('text-danger fw-bold')
            return false
        }else{
            $('#q'+(totalQuestions)+'answer'+i).removeClass('border border-danger')
            $('#aLabelTitle').removeClass('text-danger fw-bold')
            return true
        }
    }
    validateAnswerScore(score, totalQuestions, i){

        if(score == '' || typeof(score) == 'undefined' || isNaN(score)){
            $('#q'+(totalQuestions)+'score'+i).addClass('border border-danger')
            $('#aScoreTitle').addClass('text-danger fw-bold')
            return false
        }else{
            $('#q'+(totalQuestions)+'score'+i).removeClass('border border-danger')
            $('#aScoreTitle').removeClass('text-danger fw-bold')
            return true
        }
    }

    // validate report
    ruleValueIsValid(n){
        if( n != "" && !isNaN(n) ){
            // valid
            $('#ruleValue').removeClass('border border-danger')
            $('#ruleValueLabel').removeClass('text-danger fw-bold')
            return true
         }else{
            //  non valid
             $('#ruleValue').addClass('border border-danger')
             $('#ruleValueLabel').addClass('text-danger fw-bold')
             return false
         }
    }
    validateReport(totalRules){
        for (let index = 0; index < totalRules; index++) {
            const ruleComment = $('#ruleComment'+totalRules)
            if (ruleComment.val() == ''){
                ruleComment.addClass('border border-danger')
                return false
            }else{
                ruleComment.removeClass('border border-danger')
                return true
            }
        }
    }

    // add question
    questionBlockHeader(counter, answersNumber){
        return `
        <div class="col-sm">
            <h3>Question ${counter}</h3>
            
            <input type="hidden"name="answersNumber${counter}" id="answersNumber${counter}" value="${answersNumber}" />
            <div class="row mb-1">
                <div class="col-sm"> story: </div>
                <div class="col-sm"> 
                    <textarea class="form-control" name="questionStory${counter}" id="questionStory${counter}" rows="3"></textarea>
                </div>
            </div>
            <div class="row mb-1">
                <div class="col-sm" id="questionLabelLabel${counter}"> label: </div>
                <div class="col-sm"> 
                    <input type="text" class="form-control" name="questionLabel${counter}" id="questionLabel${counter}" />
                </div>
            </div>
            <div class="row mb-1">
                <div class="col-sm"> helper: </div>
                <div class="col-sm"> 
                    <input type="text" class="form-control" name="questionHelper${counter}" id="questionHelper${counter}" />
                </div>
            </div>
            <div class="row mb-3">
                <div class="col-sm"> answers: </div>
                <div class="col-sm">
                    <div class="row">
                        <div class="col-sm-5" id="aLabelTitle">Label</div>
                        <div class="col-sm-2" id="aScoreTitle">Score</div>
                        <div class="col-sm-5" id="aCommentTitle">Comment</div>
                    </div>
        `
    }
    addCheckboxAnswers(counter, checkboxesNumber){
        const target = document.querySelector("#formContener"),
              node = document.createElement("div")
        node.classList.add("row", "align-items-end", "justify-content-center", "mb-1")
        node.id = "question"+counter

        let contener = this.questionBlockHeader(counter, checkboxesNumber)
         
        for (let index = 0; index < checkboxesNumber; index++) {
            contener += `
            <div class="row">
                <div class="col-sm-5">
                <div class="form-check">
                    <input class="form-check-input" type="checkbox" value="" id="q${counter}checkbox${index}" disabled>
                    <label class="form-check-label" for="q${counter}checkbox${index}">
                        <input type="text" class="form-control" name="q${counter}answer${index}" id="q${counter}answer${index}" />
                    </label>
                </div>
                </div>
                <div class="col-sm-2"><input type="text" class="form-control" name="q${counter}score${index}" id="q${counter}score${index}" /></div>
                <div class="col-sm-5"><input type="text" class="form-control" name="q${counter}comment${index}" id="q${counter}comment${index}" /></div>
            </div>
            
            `
        }
        contener += `
                </div>
            </div>
            <input type="hidden"name="questionType${counter}" id="questionType${counter}" value="1" />
        </div>`
        node.innerHTML = contener
        target.appendChild(node)
    }
    addRadioAnswers(counter, radiosNumber){
        const target = document.querySelector("#formContener"),
              node = document.createElement("div")
        node.classList.add("row", "align-items-end", "justify-content-center", "mb-1")
        node.id = "question"+counter

        let contener = this.questionBlockHeader(counter, radiosNumber)

        for (let index = 0; index < radiosNumber; index++) {
            contener += `
            <div class="row">
                <div class="col-sm-5">
                <div class="form-check">
                    <input class="form-check-input" type="radio" value="" id="q${counter}radio${index}" disabled>
                    <label class="form-check-label" for="q${counter}radio${index}">
                        <input type="text" class="form-control" name="q${counter}answer${index}" id="q${counter}answer${index}" />
                    </label>
                </div>
                </div>
                <div class="col-sm-2"><input type="text" class="form-control" name="q${counter}score${index}" id="q${counter}score${index}" /></div>
                <div class="col-sm-5"><input type="text" class="form-control" name="q${counter}comment${index}" id="q${counter}comment${index}" /></div>
            </div>
            
            `
        }
        contener += `
                </div>
            </div>
            <input type="hidden"name="questionType${counter}" id="questionType${counter}" value="2" />
        </div>`
        node.innerHTML = contener
        target.appendChild(node)
    }
    addFreeAnswer(counter, keywordsNumber){
        const target = document.querySelector("#formContener"),
              node = document.createElement("div")
        node.classList.add("row", "align-items-end", "justify-content-center", "mb-1")
        node.id = "question"+counter

        let contener = this.questionBlockHeader(counter, keywordsNumber)

        for (let index = 0; index < keywordsNumber; index++) {
            contener += `
            <div class="row">
                <div class="col-sm-5">
                <div class="form-check">
                    <label class="form-check-label" for="q${counter}radio${index}">
                        <input type="text" class="form-control" name="q${counter}answer${index}" id="q${counter}answer${index}" />
                    </label>
                </div>
                </div>
                <div class="col-sm-2"><input type="text" class="form-control" name="q${counter}score${index}" id="q${counter}score${index}" /></div>
                <div class="col-sm-5"><input type="text" class="form-control" name="q${counter}comment${index}" id="q${counter}comment${index}" /></div>
            </div>
            
            `
        }
        contener += `
                </div>
            </div>
            <input type="hidden"name="questionType${counter}" id="questionType${counter}" value="3" />
        </div>`
        node.innerHTML = contener
        target.appendChild(node)
    }

    // add rule
    addOperatorRule(counter, ruleValue, ruleMessage){
        const target = document.querySelector("#formContenerRules"),
        node = document.createElement("div")
        node.classList.add("row", "align-items-end", "justify-content-center", "mb-3")
        node.id = "rule"+counter

        let contener =`
        <h3>Rule ${counter}</h3>
        <input type="hidden" name="ruleValue${counter}" id="ruleValue${counter}" value="${ruleValue}" />
        <input type="hidden" name="ruleOperator${counter}" id="ruleOperator${counter}" value="${ruleMessage}" />

        <div class="col-sm">
            <span class="form-label fw-bolder fst-italic fs-6"><small><small><small>When the survey's score results are ${ruleMessage} ${ruleValue}%</small></small></small></span>
            <textarea class="form-control" name="ruleComment${counter}" id="ruleComment${counter}" rows="3"></textarea>
        </div>
        `

        node.innerHTML = contener
        target.appendChild(node)
    }

    // not used...to post form via ajax
    // serialize (data) {
    //     let obj = {};
    //     for (let [key, value] of data) {
    //         if (obj[key] !== undefined) {
    //             if (!Array.isArray(obj[key])) {
    //                 obj[key] = [obj[key]]
    //             }
    //             obj[key].push(value)
    //         } else {
    //             obj[key] = value
    //         }
    //     }
    //     return obj
    // }

}