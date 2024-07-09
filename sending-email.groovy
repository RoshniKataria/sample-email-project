pipeline{
    agent any
    parameters{
        string(
            name: 'RecipientEmailId',
            description: 'Enter the emails to which you want to send the request'
        )
        string(
            name: 'CafeName',
            description: 'Enter the cafe name with which you want to collaborate'
        )
    }
    stages{
        stage('Email'){
            steps{
                script{
                    def emailBody = """
                        Hi ${params.CafeName} Team, 

                        I hope this email finds you well! My name is Roshni Kataria, and I'm an Instagram influencer specialising in café/restaurant reviews. I recently came across your profile and was impressed by the vibrant atmosphere and cozy vibes you bring to the city. 

                        I'd love the opportunity to visit your place, experience its unique ambiance, and share my review with my 15k+ engaged viewers on Instagram. Would you be open to collaborating on a feature? 

                        Looking forward to the possibility of working together! 
                        
                        Best regards, 
                        Roshni Kataria 
                        Phone Number: +91 9503021415 

                        Instagram: https://www.instagram.com/roshnii_kataria?igsh=ZDBqMGIydmp3YWJk&utm_source=qr
                    """
                    def emailSubject = "Collaboration Opportunity: Reviewing Your Cafe: ${params.CafeName} on Instagram!"
                    emailext(
                        to: "${params.RecipientEmailId}",
                        subject: "${emailSubject}",
                        body: "${emailBody}",
                        mimeType: 'text/plain'
                    )
                }
            }
        }
    }
}