class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-654f46bc058ea34cb4bf1373b439b902cf699b1f/example-go.tar.gz"
  version "654f46bc058ea34cb4bf1373b439b902cf699b1f"

  def install
    bin.install "example-go"
    man1.install Dir["*.1"]
    bash_completion.install "example-go.bash" => "example-go"
    zsh_completion.install "example-go.zsh" => "_example-go"
    fish_completion.install "example-go.fish"
  end
end
