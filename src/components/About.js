import React from 'react';
import './About.css';
import human1 from './human1.jpg';
import human2 from './human2.jpg';
import human3 from './human3.jpg';
import human4 from './human4.jpg';

const About = () => {
    return (
        <div>
            <h2>About Us</h2>
            <p>We are a leading travel booking platform dedicated to providing the best travel experiences.</p>
        
            <div className="team-members">
                <div className="team-member">
                    <a href="https://www.linkedin.com/in/prathima-dodda/" target="_blank" rel="noopener noreferrer">
                        <img src={human1} alt="Prathima Dodda" />
                    </a>
                    <p>Prathima Dodda</p>
                </div>
                <div className="team-member">
                <a href="https://www.linkedin.com/in/sai-preethi123/" target="_blank" rel="noopener noreferrer">
                <img src={human2} alt="Sai Preethi Kota" />
                    </a>
                    <p>Sai Preethi Kota</p>
                </div>
                <div className="team-member">
                <a href="https://www.linkedin.com/in/kopparla-varshini/" target="_blank" rel="noopener noreferrer">
                <img src={human3} alt="Varshini Kopparla" />
                    </a>
                    
                    <p>Varshini Kopparla</p>
                </div>
                <div className="team-member">
                <a href="https://www.linkedin.com/in/karthik-karnam/" target="_blank" rel="noopener noreferrer">
                <img src={human4} alt="Karthik Karnam" />
                    </a>
                    
                    <p>Karthik Karnam</p>
                </div>
            </div>
        </div>
    );
};

export default About;
