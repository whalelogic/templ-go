package models


type Experience struct {
        CompanyName string
        Role        string
        StartDate   string
        EndDate     string
        Description string
}

type Skill struct {
        Name        string
        Proficiency string
}

type Language struct {
        Name        string
        Proficiency string
}

type Interest struct {
        Name string
}

type Education struct {
        SchoolName   string
        Degree       string
        FieldOfStudy string
}

type Profile struct {
        Name        string
        Title       string
        Bio         string
        AvatarURL   string
        Educations  []Education
        Experiences []Experience
        Skills      []Skill
        Languages   []Language
        Interests   []Interest
}


var MyProfile = Profile{
        Name:      "Keith Thomson",
        Title:     "Software Engineer - Cloud Solutions Architect",
        Bio:       "Passionate developer with experience in Go and web development.",
        AvatarURL: "https://gravatar.com/avatar/24289225bb8abcf8cdd1fe4d5d20db4d?s=400&d=robohash&r=x",
        Educations: []Education{
                {
                        SchoolName:   "Southern Connecticut State University",
                        Degree:       "Bachelor of Science",
                        FieldOfStudy: "Computer Science - Information Systems",
                },
                {
                        SchoolName:   "Connecticut State University",
                        Degree:       "Associate of Science",
                        FieldOfStudy: "Cloud Computing & Virtualization",
                },
        },
        Experiences: []Experience{
                {
                        CompanyName: "WhalerAPI",
                        Role:        "Founder & Lead Developer",
                        StartDate:   "Mar 2025",
                        EndDate:     "Present",
                        Description: "Designing Cloud Solutions & Developing scalable backend services using Go and Python.",
                },
                {
                        CompanyName: "PennyBorne Development LLC",
                        Role:        "Product QA Lead - Content Management System",
                        StartDate:   "Jan 2022",
                        EndDate:     "Current",
                        Description: "Developing scalable backend services using Go.",
                },
                {
                        CompanyName: "US Army",
                        Role:        "Infantryman",
                        StartDate:   "Sep 2007",
                        EndDate:     "Sep 2011",
                        Description: "Mechanized Infantry, Squad Leader, Served in Iraq 2008-2010",
                },
        },
        Skills: []Skill{
                {Name: "Go", Proficiency: "Expert"},
                {Name: "Python", Proficiency: "Advanced"},
                {Name: "Docker", Proficiency: "Advanced"},
                {Name: "Kubernetes", Proficiency: "Intermediate"},
        },
        Languages: []Language{
                {Name: "English", Proficiency: "Native"},
                {Name: "Spanish", Proficiency: "Conversational"},
                {Name: "German", Proficiency: "Basic"},
        },
        Interests: []Interest{
                {Name: "Open Source Contribution"},
                {Name: "Cloud Computing"},
                {Name: "Traveling"},
        },
}


