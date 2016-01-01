## Showcase Projects

Display your projects in a gallery


### How to add project

1. Fork this repo
2. Clone your fork with:
   ```
   git clone git@github.com:<your-username>/showcase-projects.git
   ```
3. Update `projects.yaml` file and include your project details. Instructions [below](#instructions).
4. Commit your work with:
   ```
   git commit -m "Add project: <project-name>"
   ```
   If you have more than one project, then commit them separately.
5. Push your changes with:
   ```
   git push origin master
   ```


### Update `projects.yaml` <a name="instructions"></a>

You must enter the following details in YAML format:

- **project**: Project name.
- **description**: One line description of your project, possibly in less than 80 chars.
- **website-link**: Link to the application. If it is a website then link to the site, or if it is an android app, link to Play Store.
- **github-link**: Link to the GitHub source repo.
- **tags**: A list of tags of the technologies used. Keep each tag in lowercase.
  e.g. `[python, html, css, javascript]`
- **creator**: Creator name.
- **creator-link**: Link to the creator. You can link to your website, blog, Twitter page or even your GitHub profile.

**Note**: **website-link** is optional.


### License

See the `LICENCE` file.
