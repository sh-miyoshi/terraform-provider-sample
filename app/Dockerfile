FROM ruby:2.7-slim

WORKDIR /app
COPY Gemfile Gemfile
COPY Gemfile.lock Gemfile.lock
RUN bundle install
COPY main.rb main.rb

CMD ["ruby", "main.rb", "-o", "0.0.0.0"]
