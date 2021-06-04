FROM ruby:2.7
RUN gem install bundler
WORKDIR /myapp
COPY Gemfile /myapp/Gemfile
RUN bundle install
COPY name.rb /myapp

EXPOSE 8000
CMD [ "ruby", "name.rb"]
