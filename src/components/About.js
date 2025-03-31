import React from 'react';
import './About.css';
import human1 from './human1.jpg';
import human2 from './human2.jpg';
import human3 from './human3.jpg';
import human4 from './human4.jpg';

const About = () => {
  return (
    <div className="about-container">
      <section className="about-hero">
        <div className="hero-content">
          <h1>About Our Journey</h1>
          <p>
            We are a forward-thinking travel booking platform dedicated to crafting memorable travel experiences in real time.
          </p>
        </div>
      </section>

      <section className="about-info">
        <h2>Our Vision &amp; Values</h2>
        <p>
          Leveraging cutting-edge technology, we provide seamless travel solutions—from hotel accommodations to custom travel packages—ensuring every journey is exceptional.
        </p>
      </section>

      <section className="team-section">
        <h2>Meet Our Team</h2>
        <div className="team-grid">
          <div className="team-member">
            <a href="https://www.linkedin.com/in/prathima-dodda/" target="_blank" rel="noopener noreferrer">
              <img src={human1} alt="Prathima Dodda" />
            </a>
            <span>Prathima Dodda</span>
          </div>
          <div className="team-member">
            <a href="https://www.linkedin.com/in/sai-preethi123/" target="_blank" rel="noopener noreferrer">
              <img src={human2} alt="Sai Preethi Kota" />
            </a>
            <span>Sai Preethi Kota</span>
          </div>
          <div className="team-member">
            <a href="https://www.linkedin.com/in/kopparla-varshini/" target="_blank" rel="noopener noreferrer">
              <img src={human3} alt="Varshini Kopparla" />
            </a>
            <span>Varshini Kopparla</span>
          </div>
          <div className="team-member">
            <a href="https://www.linkedin.com/in/karthik-karnam/" target="_blank" rel="noopener noreferrer">
              <img src={human4} alt="Karthik Karnam" />
            </a>
            <span>Karthik Karnam</span>
          </div>
        </div>
      </section>
    </div>
  );
};

export default About;
