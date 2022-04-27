class Survey_UI{
    // Check if value already exist in array
    // Return bool
    valueAlreadyExist(arr=[], value=''){
        let exist
        return (arr.includes(value)) ? exist = true : exist = false         
    }
    // Return the list of all names of the element class
    getAllElementsName(classname){
        const myClass = this
        let nameList = []
        $('.'+classname).each(function() {
            let elementName = $(this)[0].name
      
            if(!myClass.valueAlreadyExist(nameList, elementName)){
                nameList.push(elementName)
            }            
        })
        console.log(nameList);
        return nameList
    }
    // Check if element got at least one option(checkbox|radio) checked
    // Return bool
    elementGotAtLeastOneChecked(inputName){
        const elementChecked = $('input[name="'+inputName+'"]:checked'),
              elements = $('input[name="'+inputName+'"]')

        
        if (elementChecked.length > 0){
            // element.classList.remove('border','border-danger')
            this.highlightElements(elements, false)
            return true
        }else{
            // element.classList.add('border','border-danger')
            this.highlightElements(elements, true)
            return false
        }
    }
    highlightElements(elements, on){
        
        for (let index = 0; index < elements.length; index++) {
            const element = elements[index];

            if ( on == true ){
                element.classList.add('border','border-danger')
                element.nextElementSibling.classList.add('fw-bold','text-danger')
            } else {
                element.classList.remove('border','border-danger')
                element.nextElementSibling.classList.remove('fw-bold','text-danger')
                
            }

        }

    }

    elementIsNotEmpty(inputName){
        const element = $('input[name="'+inputName+'"]')[0]

        if (element.value != ''){
            element.classList.remove('border','border-danger')
            return true
        }else{
            element.classList.add('border','border-danger')
            return false
        }
        
    }
    validateSurveyForm(classnameList) {
        const myClass = this
        let  allValidations = []

        // Loop over all classname
        for (let index = 0; index < classnameList.length; index++) {
            const oneClass = classnameList[index], // element is a 
                  nameList = myClass.getAllElementsName(oneClass)

            // When got element
            if (nameList.length > 0){
                // Store all same type element isValid
                let listValidation = []
                // Loop over all input name
                for (let i = 0; i < nameList.length; i++) {
                    const element = nameList[i],
                          elementType = $('input[name="'+element+'"]')[0].type
                    let oneValidation
                    switch (elementType) {
                        case 'text':
                            oneValidation = myClass.elementIsNotEmpty(element)
                            break;                        
                        case 'radio':
                            oneValidation = myClass.elementGotAtLeastOneChecked(element)
                            break;                                
                        case 'checkbox':
                            oneValidation = myClass.elementGotAtLeastOneChecked(element)
                            break;                    
                        default:
                            break;
                    }
                    listValidation.push(oneValidation)
                } 
                if (listValidation.includes(false)){
                    allValidations.push(false)
                }else{
                    allValidations.push(true)
                }
            }

        }

        if (allValidations.includes(false)){
           return false
        }else{
            return true
        }

    }

}