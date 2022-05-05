const UI = new SurveyReport_UI()
// Load Google charts
google.charts.load("current", {packages:['corechart']});

let questionNumber = 1,
    arrayNumber = 0

// Loop over the questionnaire
for (let i = 0; i < questionnaire.length; i++) {
    const question = questionnaire[i],
          title = 'Question'+questionNumber+' - '+question.Label,
          targetID = 'chartGoogle'+questionNumber
    // console.log(title)
    // console.log(targetID)
    let data = [], header = []
    header.push("Element", title, { role: "style" } )
    data.push(header)
    // Loop over each answers of each question
    for (let j = 0; j < question.Answers.length; j++) {
        let groupData = []
        const element = question.Answers[j]
        let totalAnswered = 0
        // Loop over the evaluations
        for (let x = 0; x < evaluations.length; x++) {
            const evaluation = evaluations[x];
            // console.log(evaluation.Answers[arrayNumber].Response);
            // Loop over each answered question of each evaluation
            for (let y = 0; y < evaluation.Answers[arrayNumber].Response.length; y++) {
                const answered = evaluation.Answers[arrayNumber].Response[y];
                // console.log(answered);
                if (element.Label == answered || answered.includes(element.Label)){
                    totalAnswered += 1
                }
            }
        }
        
        const percentage = totalAnswered/evaluations.length*100
        // console.log('label element: '+element.Label+' ('+percentage+' %)')
        let colorColumn 
        switch (true) {
            case percentage>=50:
                colorColumn = '#42f578' // green
                break;
            case percentage<50:
                colorColumn = '#42ddf5' // light blue
                break;
            case percentage<25:
                colorColumn = '#a6bcbf' // light grey
                break;
        }
        groupData.push(element.Label, percentage, colorColumn)
        data.push(groupData)
    }

    // Create Google charts
    google.charts.setOnLoadCallback(
        function() {
            UI.drawChart(data, title, targetID)
        }
      )

    questionNumber++
    arrayNumber++
}

window.addEventListener('DOMContentLoaded', (event) => { 
    $('#printReport').bind('click', () => {
        UI.CreatePDFfromHTML()
    })    
})
